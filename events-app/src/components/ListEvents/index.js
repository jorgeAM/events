import React, { useEffect, useState } from "react";
import Event from "../Event";

const ListEvents = () => {
  const [events, setEvents] = useState([]);

  useEffect(() => {
    const fetchEvents = async () => {
      const rest = await fetch(process.env.REACT_APP_EVENT_SERVICE);
      const json = await rest.json();
      setEvents(json);
    };

    fetchEvents();
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
