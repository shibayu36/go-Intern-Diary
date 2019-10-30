import React, { useState } from "react";
import gql from "graphql-tag";
import { useMutation } from "@apollo/react-hooks";
import { CreateDiary, CreateDiaryVariables } from "./__generated__/CreateDiary";
import { Redirect } from "react-router";

const createDiaryMutation = gql`
  mutation CreateDiary($name: String!) {
    createDiary(name: $name) {
      id
      name
    }
  }
`;

interface DiaryCreateFormState {
  name: string
}

export const DiaryCreateForm: React.FunctionComponent = () => {
  const [name, setName] = useState('');
  const [createDiary, { error: error, data: createdDiary }] = useMutation<CreateDiary, CreateDiaryVariables>(createDiaryMutation);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const input = e.currentTarget;
    switch (input.name) {
      case "name":
        setName(input.value)
        break;
    }
  }

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    createDiary({
      variables: { name: name }
    })
  }

  if (createdDiary) return <Redirect to="/diaries" />

  return <form className="DiaryCreateForm" onSubmit={handleSubmit}>
    {error && <p>{error.message}</p>}
    <div><label>ダイアリー名: <input
      type="text"
      name="name"
      value={name}
      onChange={handleChange} /></label></div>
    <input type="submit" value="作成" />
  </form >

}
