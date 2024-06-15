import axios from "axios";
import queryString from "query-string";
import { useEffect, useState } from "react";
import TracksFound from "../components/TracksFound";
import TracksSelected from "../components/TracksSelected";

import "../styles/components/create-playlist.css";
import "../styles/pages/playlist.css";

const MAX_SEED_TRACKS = 5;

export default function Playlist() {
  const [accessToken, setAccessToken] = useState();
  const [tracksFound, setTracksFound] = useState([]);
  const [tracksSelected, setTracksSelected] = useState([]);
  const [tracksSelectedInfo, setTracksSelectedInfo] = useState({});
  const [searchTerm, setSearchTerm] = useState("");

  useEffect(() => {
    const token = localStorage.getItem("spotify-utils:access_token");
    setAccessToken(token);
  }, [accessToken, setAccessToken]);

  useEffect(() => {
    const searchForTrack = async (text) => {
      if (!(text?.length >= 3)) {
        setTracksFound([]);
        return;
      }
      
      const accessToken = localStorage.getItem("spotify-utils:accessToken");
      const querystring = queryString.stringify({
        q: text,
        type: "track",
        limit: 10,
      });
      const url = encodeURI(`http://localhost:5555/search?${querystring}`);
      const response = await axios.get(url, {
        headers: {
          Authorization: "Bearer " + accessToken,
        },
      });

      setTracksFound(response?.data?.tracks?.items || []);

      // const r = await axios.get("http://localhost:5555/playlist/" , {
      //   headers: {
      //     Authorization: "Bearer " + accessToken,
      //   },
      // });

      // console.log(r.data)
    };
    const timeOutId = setTimeout(() => searchForTrack(searchTerm), 300);
    return () => clearTimeout(timeOutId);
  }, [searchTerm]);

  const toggleTrackSelection = (trackId) => {
    if (tracksSelected.includes(trackId)) {
      setTracksSelected(tracksSelected.filter((id) => id !== trackId));
      console.log(tracksSelected);
      return;
    }

    if (tracksSelected.length === MAX_SEED_TRACKS) {
      alert("You can only select up to 5 tracks");
      return;
    }

    setTracksSelected([...tracksSelected, trackId]);
    const track = tracksFound.find(({ id }) => id === trackId);
    console.log(track);
    setTracksSelectedInfo({
      ...tracksSelectedInfo,
      [track.id]: {
        name: track.name,
        artists: track.artists.map(({ name }) => name).join(", "),
        album: {
          name: track.album.name,
          cover: track.album.images[0].url,
        },
      },
    });
  };

  const handlePlaylistCreation = async (event) => {
    event.preventDefault();
    if (!tracksSelected.length) {
      alert("Please select a track");
      return;
    }

    const playlistName = event.target[0].value;
    if (!playlistName.length) {
      alert("Please type a name for the playlist");
      return;
    }

    const userId = localStorage.getItem("spotify-utils:userId");
    const accessToken = localStorage.getItem("spotify-utils:accessToken");

    console.log(accessToken, playlistName, tracksSelected.join(","), userId);

    const response = await axios.post(
      `http://localhost:5555/${userId}/playlist/new`,
      { playlistName, trackIds: tracksSelected.join(",") },
      {
        headers: {
          Authorization: "Bearer " + accessToken,
        },
      }
    );

    console.log(response.data);
  };

  return (
    <div>
      <div className="search-bar">
        <input
          className="search-input"
          name="trackToSearch"
          onChange={(e) => setSearchTerm(e.target.value)}
          placeholder="Search"
          type="text"
        />
      </div>
      <div className="main-panel">
        <TracksFound
          toggleTrackSelection={toggleTrackSelection}
          tracksFound={tracksFound}
          tracksSelected={tracksSelected}
        />

        <TracksSelected
          tracksSelected={tracksSelected}
          tracksSelectedInfo={tracksSelectedInfo}
          toggleTrackSelection={toggleTrackSelection}
        />
      </div>

      <div className="create-playlist-bar">
        <form onSubmit={handlePlaylistCreation}>
          <input
            className="playlist-name-input"
            placeholder="Nome da playlist"
            type="text"
            name="name"
          />
          <input
            className="create-playlist-button"
            type="submit"
            value="Create"
          />
        </form>
      </div>
    </div>
  );
}
