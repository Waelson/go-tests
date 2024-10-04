import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: 10,
    duration: '600s',
};

export default function () {
    const url = 'http://localhost:8080/users';
    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const res = http.get(url, params);

    check(res, {
        'status was 200': (r) => r.status === 200,
    });

    sleep(1);
}
