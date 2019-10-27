import React from "react";
import { DiaryCreateForm } from "../../components/diary_create_form";

export const DiariesCreate: React.FunctionComponent = () => {
  return <div className="DiariesCreate">
    <h1>ダイアリー作成</h1>
    <DiaryCreateForm />
  </div>
}
