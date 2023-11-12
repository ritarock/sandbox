import { useParams } from "react-router-dom";

export default function NotFoundPage() {
  const { "*": paths } = useParams();
  return <p>{paths} is NotFound</p>;
}
