/** @jsx jsx */
import { css, jsx } from "@emotion/core";

const style = css`
  display: flex;
  width: 100%;
  justify-content: space-around;
`;

const EventDetail = ({ startAt, available }) => (
  <div css={style}>
    <p>Empieza: {startAt}</p>
    {available ? <p>{"disponible"}</p> : <p>{"no disponible"}</p>}
  </div>
);

export default EventDetail;
