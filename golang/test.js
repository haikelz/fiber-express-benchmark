import { check, sleep } from "k6";
import http from "k6/http";

const env = {
  vus: process.env.K6_VUS,
  duration: process.env.K6_DURATION,
  projectID: process.env.K6_PROJECT_ID,
  apiUrl: process.env.K6_API_URL,
};

export const options = {
  vus: env.vus,
  duration: env.duration,
  cloud: {
    projectID: env.projectID,
  },
};

export default function () {
  const res = http.get(`${env.apiUrl}/api/all`);
  check(res, {
    "status was 200": (r) => r.status == 200,
  });
  sleep(1);
}
