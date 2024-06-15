import { useEffect } from "react";

const randSeq = ({ length }) => {
  const characters = "abcdefghijklmnopqrstuvwxyz";
  let result = " ";
  const charactersLength = characters.length;
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }

  return result;
};

const CLIENT_ID = ""
const REDIRECT_URI = "http://localhost:3001/"

const redirectToAuthPage = () => {
  const scope =
    "user-read-private user-read-email playlist-read-private playlist-modify-public playlist-modify-private";
  const state = randSeq(16);

  const url = new URL("https://accounts.spotify.com/authorize");
  const params = new URLSearchParams({
    "response_type": "code",
    "client_id": CLIENT_ID,
    "scope": scope,
    "redirect_uri": REDIRECT_URI,
    "state": state
  });
  url.search = params.toString();

  console.log(url.toString());
  window.location.replace(url.toString());
};

export default function Auth() {
  useEffect(() => {
    console.log("Authing...");
    redirectToAuthPage();
  }, []);

  return <div>Auth</div>;
}
