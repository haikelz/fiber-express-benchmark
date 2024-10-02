import { check, sleep } from "k6";
import http from "k6/http";

export const options = {
  vus: process.env.K6_VUS,
  duration: process.env.K6_DURATION,
  cloud: {
    projectID: process.env.K6_PROJECT_ID,
  },
};

export default function () {
  const res = http.get(`${process.env.K6_API_URL}/api/all`);
  check(res, {
    "status was 200": (r) => r.status == 200,
  });
  sleep(1);
}
