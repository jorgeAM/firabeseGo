import firebase from "firebase/app";
import "firebase/auth";
import { ChakraProvider } from "@chakra-ui/react";
import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";

const firebaseConfig = {
  apiKey: "AIzaSyBPKkUdg_bOINo99Lyq4G7POLbOtAN1L5o",
  authDomain: "golang-authentication.firebaseapp.com",
  projectId: "golang-authentication",
  storageBucket: "golang-authentication.appspot.com",
  messagingSenderId: "947172532815",
  appId: "1:947172532815:web:28a4935bd9f36baafbc03a",
};

firebase.initializeApp(firebaseConfig);

ReactDOM.render(
  <React.StrictMode>
    <ChakraProvider>
      <App />
    </ChakraProvider>
  </React.StrictMode>,
  document.getElementById("root")
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
