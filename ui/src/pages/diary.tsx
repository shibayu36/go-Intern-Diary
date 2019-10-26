import React from "react";
import gql from "graphql-tag";
import { useQuery } from "@apollo/react-hooks";
import { GetDiary, GetDiaryVariables } from "./__generated__/GetDiary";
import { RouteComponentProps, useParams } from "react-router";

const getDiaryQuery = gql`
  query GetDiary($diaryId: ID!) {
    diary(diaryId: $diaryId) {
      id
      name
    }
  }
`

interface RouteProps {
  diaryId: string;
}

export const Diary: React.FunctionComponent = () => {
  const { diaryId } = useParams<RouteProps>();
  const { loading, error, data } = useQuery<GetDiary, GetDiaryVariables>(
    getDiaryQuery,
    {
      variables: { diaryId: diaryId }
    }
  );

  if (loading) return <p>Loading...</p>;
  if (error) return <p>{error.message}</p>;

  return <div className="Diary">
    <h1>{data!.diary.name}</h1>
  </div>;
}
