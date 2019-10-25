import React from "react";
import { NavLink } from "react-router-dom";
import gql from "graphql-tag";
import { useQuery } from '@apollo/react-hooks'

const visitorQuery = gql`
  query GetVisitor {
    visitor {
      id
      name
    }
  }
`;

export const GlobalHeader = () => {
  const { loading, error, data } = useQuery(visitorQuery);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error</p>;
  console.log(error)

  return <header className="GlobalHeader">
    ようこそ、{data.visitor.name}さん
    <nav>
      <ul>
        <li><NavLink to="/">トップ</NavLink></li>
      </ul>
    </nav>
  </header>
};
