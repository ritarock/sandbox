import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  thresholds: {
    http_req_duration: ["p(90)<400", "p(95)<800", "p(99)<2000"],
  },
};

export default function () {
  http.get("https://test-api.k6.io/public/crocodiles/1/");
  sleep(1);
}
