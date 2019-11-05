import React from 'react';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import { GetMyDiaries } from './__generated__/GetMyDiaries';
import { Link } from 'react-router-dom';

export const getMyDiariesQuery = gql`
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
`;

export const Diaries: React.FunctionComponent = () => {
  const { loading, error, data } = useQuery<GetMyDiaries>(getMyDiariesQuery);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>{error.message}</p>;

  const diaries = data!.visitor.diaries;

  return (
    <div className="Diaries">
      <h1>{data!.visitor.name}のダイアリー一覧</h1>
      {diaries.map((diary: any) => (
        <Link key={diary.id} to={`/diaries/${diary.id}`}>
          <div>
            <p>{diary.name}</p>
          </div>
        </Link>
      ))}

      <div>
        <Link to="/diaries/create">新しいダイアリーを作成</Link>
      </div>
    </div>
  );
};
