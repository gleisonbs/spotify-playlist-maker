import axios from "axios";
import { useEffect, useState } from "react";
import { Link, useSearchParams } from "react-router-dom";
import env from "react-dotenv";

const REDIRECT_URI = env.REDIRECT_URI;
const CLIENT_ID = env.CLIENT_ID;

const makeUserRequest = async ({ accessToken }) => {
  const response = await axios.get("https://api.spotify.com/v1/me", {
    headers: { Authorization: "Bearer " + accessToken },
  });
  if (response.status === 200) {
    const { id } = response.data;
    localStorage.setItem("spotify-utils:user_id", id);
  }
};

const getToken = async ({ code, setAccessToken }) => {
  let codeVerifier = localStorage.getItem('code_verifier');

  const payload = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: new URLSearchParams({
      client_id: CLIENT_ID,
      grant_type: 'authorization_code',
      code,
      redirect_uri: REDIRECT_URI,
      code_verifier: codeVerifier,
    }),
  }

  console.log({
    client_id: CLIENT_ID,
    grant_type: 'authorization_code',
    code,
    redirect_uri: REDIRECT_URI,
    code_verifier: codeVerifier,
  })

  const url = "https://accounts.spotify.com/api/token";
  const body = await fetch(url, payload);
  const response = await body.json();

  const { access_token: accessToken } = response
  setAccessToken(accessToken);
  localStorage.setItem("spotify-utils:access_token", accessToken);
}

export default function Main() {
  const [searchParams] = useSearchParams();
  const [accessToken, setAccessToken] = useState();

  useEffect(() => {
    const code = searchParams.get("code");
    
    if (!accessToken) {
      getToken({ code, setAccessToken });
    } else {
      makeUserRequest({ accessToken });
    }
  }, [accessToken, searchParams]);

  return (
    <div>
      Main
      <br />
      {accessToken ? (
        <Link to={"/playlist"}>Make playlist</Link>
      ) : (
        <Link to={"/auth"}>Authenticate</Link>
      )}
    </div>
  );
}
