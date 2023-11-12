import { useParams } from "react-router-dom";

export default function SearchPage() {
  const { "*": keywords } = useParams();
  // このあと String#split メソッドで / を分解して使う場面が多い
  return <p>{keywords} page</p>;
}
