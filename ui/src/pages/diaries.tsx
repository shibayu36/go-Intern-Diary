import React from "react";
import gql from "graphql-tag";
import { useQuery } from "@apollo/react-hooks";
import { GetMyDiaries } from "./__generated__/GetMyDiaries"

const MyDiariesQuery = gql`
  query GetMyDiaries {
    visitor {
      id
      name
      diaries {
        id
        name
      }
    }
  }
`

export const Diaries: React.StatelessComponent = () => {
  const { loading, error, data } = useQuery<GetMyDiaries>(MyDiariesQuery);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>{error.message}</p>;

  const diaries = data!.visitor.diaries;

  return <div className="Diaries">
    <h1>{data!.visitor.name}のダイアリー一覧</h1>
    {diaries.map((diary: any) => (
      <div><p>{diary.name}</p></div>
    ))}
  </div>;
}
