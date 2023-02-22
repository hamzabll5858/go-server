import http from 'k6/http';
import { sleep, check } from 'k6';
import { Counter } from 'k6/metrics';

// A simple counter for http requests

export const requests = new Counter('http_reqs');
const requestsGo = new Counter('http_reqs_go');

// you can specify stages of your test (ramp up/down patterns) through the options object
// target is the number of VUs you are aiming for

export const options = {
    discardResponseBodies: true,
    scenarios: {
        contacts: {
            executor: 'ramping-vus',
            startVUs: 50,
            stages: [
                { duration: '30s', target: 100 },
                { duration: '30s', target: 200 },
            ],
            gracefulRampDown: '30s',
        },
    },
    thresholds: {
        http_reqs: ['count <= 5000'],
    },
};

export default function () {
    // our HTTP request, note that we are saving the response to res, which can be accessed later

    const res = http.get('http://localhost:8080/users/2');

    sleep(0.1);

    const checkRes = check(res, {
        'status is 200': (r) => r.status === 200,
    });
}