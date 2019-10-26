import React from "react";
import { NavLink } from "react-router-dom";
import gql from "graphql-tag";
import { useQuery } from '@apollo/react-hooks'

const getVisitorQuery = gql`
  query GetVisitor {
    visitor {
      id
      name
    }
  }
`;

export const GlobalHeader = () => {
  const { loading, error, data } = useQuery(getVisitorQuery);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error</p>;

  return <header className="GlobalHeader">
    ようこそ、{data.visitor.name}さん
    <nav>
      <ul>
        <li><NavLink to="/">トップ</NavLink></li>
        <li><NavLink to="/diaries">マイダイアリー</NavLink></li>
      </ul>
    </nav>
  </header>
};
