import "firebase/auth";
import app from "firebase/app";
import {
  Flex,
  Input,
  FormControl,
  FormLabel,
  Box,
  Button,
} from "@chakra-ui/react";
import { useState } from "react";

const Login = () => {
  const auth = app.auth();

  const [payload, setPayload] = useState({
    email: "",
    password: "",
  });

  const onClickLogin = async (ev) => {
    ev.preventDefault();

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    };

    try {
      const res = await fetch("http://localhost:4000/login", requestOptions);
      const jsonRes = await res.json();

      const token = jsonRes?.user?.token;

      const firabaseRes = await auth.signInWithCustomToken(token);

      console.log(firabaseRes);
    } catch (err) {
      console.log("err: ", err);
    }
  };

  const onChangeInput = (ev) => {
    setPayload({
      ...payload,
      [ev.target.name]: ev.target.value,
    });
  };

  return (
    <Flex
      height={"100%"}
      justifyContent={"center"}
      alignItems={"center"}
      flexDirection={"column"}
    >
      <Box>
        <FormControl id="email">
          <FormLabel>Correo Electrónico</FormLabel>
          <Input
            type="email"
            value={payload.email}
            onChange={onChangeInput}
            name={"email"}
          />
        </FormControl>
        <FormControl id="password">
          <FormLabel>Contraseña</FormLabel>
          <Input
            type="password"
            value={payload.password}
            onChange={onChangeInput}
            name={"password"}
          />
        </FormControl>
        <Button colorScheme="blue" onClick={onClickLogin}>
          Registrarte
        </Button>
      </Box>
    </Flex>
  );
};

export default Login;
