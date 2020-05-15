import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import ListEvent from "./components/ListEvents";
import Booking from "./components/booking";

const App = () => {
  return (
    <>
      <Router>
        <Switch>
          <Route exact path="/">
            <ListEvent />
          </Route>
          <Route exact path="/events">
            <ListEvent />
          </Route>
          <Route exact path="/events/:id/booking">
            <Booking />
          </Route>
        </Switch>
      </Router>
    </>
  );
};

export default App;
