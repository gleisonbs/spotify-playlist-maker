import axios from "axios";
import { Buffer } from "buffer";
import { useEffect, useState } from "react";
import { Link, useSearchParams } from "react-router-dom";
import env from "react-dotenv";

const REDIRECT_URI = env.REDIRECT_URI;
const CLIENT_ID = env.CLIENT_ID;
const CLIENT_SECRET = env.CLIENT_SECRET;

const makeUserRequest = async ({ accessToken }) => {
  const response = await axios.get("https://api.spotify.com/v1/me", {
    headers: { Authorization: "Bearer " + accessToken },
  });
  if (response.status === 200) {
    const { id } = response.data;
    localStorage.setItem("spotify-utils:userId", id);
  }
};

const makeAuthRequestWithCode = async ({ code, setAccessToken }) => {
  const form = new URLSearchParams({
    code: code,
    redirect_uri: REDIRECT_URI,
    grant_type: "authorization_code",
  });

  const auth = Buffer.from(CLIENT_ID + ":" + CLIENT_SECRET).toString("base64");
  const authTokenUrl = "https://accounts.spotify.com/api/token";
  try {
    const response = await axios.post(authTokenUrl, form.toString(), {
      headers: {
        Authorization: "Basic " + auth,
        "Content-Type": "application/x-www-form-urlencoded",
      },
    });
    if (response?.status === 200) {
      const { access_token: accessToken } = response.data;
      setAccessToken(accessToken);
      localStorage.setItem("spotify-utils:accessToken", accessToken);
    }
    return response;
  } catch (error) {
    console.log(error);
  }
};

export default function Main() {
  const [searchParams] = useSearchParams();
  const [accessToken, setAccessToken] = useState();

  useEffect(() => {
    const code = searchParams.get("code");
    if (!accessToken) {
      makeAuthRequestWithCode({ code, setAccessToken });
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
