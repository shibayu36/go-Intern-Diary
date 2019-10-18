import React from "react";
import { NavLink } from "react-router-dom";

export const GlobalHeader = () => (
  <header className="GlobalHeader">
    <nav>
      <ul>
        <li><NavLink to="/">トップ</NavLink></li>
      </ul>
    </nav>
  </header>
);
