import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import CommentBoardPage from "./pages/CommentBoardPage";
import PostDetailPage from "./pages/PostDetailPage";
import LoginPage from "./pages/LoginPage";
import RegisterPage from "./pages/RegisterPage";
import ActivatePage from "./pages/ActivatePage";
import "./index.css";

const rootElement = ReactDOM.createRoot(document.getElementById("root"));
rootElement.render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/留言板" element={<CommentBoardPage />} />
      <Route path="/post/:id" element={<PostDetailPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/activate" element={<ActivatePage />} />
    </Routes>
  </BrowserRouter>
);
