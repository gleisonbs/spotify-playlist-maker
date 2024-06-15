import "../styles/components/tracks-found.css";
import "../styles/components/tracks-selected.css";

export default function TracksSelected({
  tracksSelected = [],
  tracksSelectedInfo = {},
  toggleTrackSelection,
}) {
  return (
    <div className="tracks-selected">
      {tracksSelected.length ? tracksSelected.map((id) => {
        return (
          <div
            className="track-found-row"
            key={id}
            onClick={(e) => toggleTrackSelection(id)}
          >
            <img
              className="track-found-row-image"
              style={{ maxWidth: "64px" }}
              src={tracksSelectedInfo[id].album.cover}
              alt="album cover"
            />
            <div className="track-found-row-text">
              <div className="track-found-row-text-song">{tracksSelectedInfo[id].name}</div>
              <div className="track-found-row-text-artist">{tracksSelectedInfo[id].artists}</div>
            </div>
          </div>
        );
      }): <div className="tracks-selected-empty">No tracks added</div>}
    </div>
  );
}
