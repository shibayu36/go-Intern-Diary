import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";

import { Index } from "./index";
import { NoMatch} from "./no_match";

export const App: React.StatelessComponent = () => (
  <BrowserRouter basename="/spa">
    <>
      <main>
        <Switch>
          <Route path="/" exact strict component={Index} />
          <Route path="*" component={NoMatch} />
        </Switch>
      </main>
    </>
  </BrowserRouter>
);
