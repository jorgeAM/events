import React, { useEffect } from "react";
import Event from "../Event";

const ListEvents = (props) => {
  const { events } = props;

  useEffect(async () => {
    const rest = await fetch(process.env.REACT_APP_EVENT_SERVICE);
    const json = await rest.json();
    console.log(json);
  }, []);

  return (
    <div>
      {events.map((event, i) => (
        <Event key={i} event={event} />
      ))}
    </div>
  );
};

export default ListEvents;
