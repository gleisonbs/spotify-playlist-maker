import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import './index.css';
import reportWebVitals from './reportWebVitals';
import Auth from './pages/Auth';
import Main from './pages/Main';
import Playlist from './pages/Playlist';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Main/>
  },
  {
    path: "/auth",
    element: <Auth/>
  },
  {
    path: "/playlist",
    element: <Playlist/>
  }
])
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
