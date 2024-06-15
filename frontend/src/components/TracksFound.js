import "../styles/components/tracks-found.css";


export default function TracksFound({
  tracksFound = [],
  tracksSelected = [],
  toggleTrackSelection,
}) {
  return (
    <div className="tracks-found">
      {tracksFound.length ? tracksFound?.map(({ album, artists, id, name }) => {
        const allArtists = artists.map(({ name }) => name).join(", ");
        return (
          <div
            className={tracksSelected.includes(id) ? 'track-found-row selected' : 'track-found-row'}
            key={id}
            onClick={(e) => {
              toggleTrackSelection(id);
            }}
          >
            <img
              className="track-found-row-image"
              style={{ maxWidth: "64px" }}
              src={album.images[0].url}
              alt="album cover"
            />
            <div className="track-found-row-text">
              <div className="track-found-row-text-song">{name}</div>
              <div className="track-found-row-text-artist">{allArtists}</div>
            </div>
          </div>
        );
      }) : <div className="tracks-found-empty">Type in the search bar to start adding songs</div>}
    </div>
  );
}
