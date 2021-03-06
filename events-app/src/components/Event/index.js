/** @jsx jsx */
import { css, jsx } from "@emotion/core";
import EventName from "./name";
import EventDetails from "./details";

const container = css`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const Event = (props) => {
  const {
    event: { id, name, available, startAt },
  } = props;

  return (
    <div css={container}>
      <EventName name={name} />
      <EventDetails id={id} available={available} startAt={startAt} />
    </div>
  );
};

export default Event;
