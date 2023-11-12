import { useSearchParams } from "react-router-dom";

export default function QueryPage() {
  const [params] = useSearchParams();

  return <p>{params.get("isbn")}</p>;
}
