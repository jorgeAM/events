import React from "react";
import ListEvent from "./components/ListEvents";

const events = [
  {
    name: "viva x el rock I",
    available: false,
    startAt: "2021-12-30T20:00:00Z",
  },
  {
    name: "viva x el rock II",
    available: false,
    startAt: "2021-12-30T20:00:00Z",
  },
  {
    name: "viva x el rock III",
    available: false,
    startAt: "2021-12-30T20:00:00Z",
  },
];

const App = () => {
  return (
    <div>
      <ListEvent events={events} />
    </div>
  );
};

export default App;
