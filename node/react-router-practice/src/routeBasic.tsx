import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
} from "react-router-dom";

import TopPage from "./TopPage";
import ArticlePage from "./ArticlePage";
import AboutPage from "./AboutPage";
import BookPage from "./bookPage";
import SearchPage from "./searchPage";
import NotFoundPage from "./NotFoundPage";
import QueryPage from "./queryPage";

// const routeBasic = createBrowserRouter([
//   {path: '/', Component: TopPage },
//   {path: '/article', Component: ArticlePage},
//   {path: '/about', Component: AboutPage}
// ])

const routeBasic = createBrowserRouter(
  createRoutesFromElements(
    <>
      <Route path="/" element={<TopPage />} />
      <Route path="/article" element={<ArticlePage />} />
      <Route path="/about" element={<AboutPage />} />
      {/* ルートパラメータ */}
      <Route path="/book/:isbn" element={<BookPage />} />
      {/* 可変長パラメータ */}
      <Route path="/search/*" element={<SearchPage />} />
      {/* NotFountPage */}
      <Route path="*" element={<NotFoundPage />} />
      {/* クエリを取得 query?isbn=111 */}
      <Route path="/query" element={<QueryPage />} />
    </>,
  ),
);

export default routeBasic;
