import http from 'k6/http';
import { sleep, check } from 'k6';
import encoding from 'k6/encoding';
import { uuidv4, randomItem } from './libs/k6-utils.js';
import { scenario } from 'k6/execution';

const vus = 500;
export const options = {
  stages: [
    { duration: '20s', target: vus },
    { duration: '20s',  target: vus },
    { duration: '20s', target: 0 },
  ],
  // vus,
  // duration: '20s',
  // iterations: 1,
};

function decodeJWT(token) {
    try {
      const parts = token.split('.');
      if (parts.length !== 3) throw new Error('Неверный формат JWT');
      return JSON.parse(String.fromCharCode.apply(null, new Uint8Array(encoding.b64decode(parts[1], 'rawstd'))));
    } catch (e) {
      console.error('Ошибка декодирования JWT:', e.message);
      return null;
    }
}

const BASE_URL = 'http://localhost';


export function setup() {
  const users = [];

  const healthRes = http.get(`${BASE_URL}/health`);
  check(healthRes, { 'service is health': (r) => r.status === 200 });
  check(healthRes, { 'database is health': (r) => r.json('database') === "OK" });

  for (let i = 0; i < vus; i++) {
    const user_id = uuidv4();
    const body = JSON.stringify({
      email: `${user_id}@email.com`,
      password: `${user_id}`
    });
    const params = {
      headers: { 'Content-Type': 'application/json' }
    };
    
    const response = http.post(`${BASE_URL}/auth/sign-up`, body, params);
    if (check(response, { 'sign up is ok': (r) => r.status === 201 })) {
      const token = response.json('token');
      const decoded = decodeJWT(token);
      if (decoded) {
        users.push({ id: decoded.user_id, token: token });
      }
    }
  }

  return users;
}


export default function (users) {
  const current_user = users[__VU - 1];
  const others = Array.from(users).filter(u => u.id !== current_user.id);
  if (others.length === 0) {
    sleep(3);
    return;
  }

  const body = JSON.stringify({
    recipient_id: randomItem(others).id, 
    amount: Math.random() * 1000,
  });
  const params = {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${current_user.token}`,
    }
  };
  const response = http.post(`${BASE_URL}/payments`, body, params);
  check(response, {'payments is ok': (r) => r.status === 201});

  sleep(3);
}
