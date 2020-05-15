/** @jsx jsx */
import { css, jsx } from "@emotion/core";
import Switch from "react-switch";
import { useState } from "react";
import moment from "moment";

const style = css`
  display: flex;
  width: 100%;
  justify-content: space-around;
  align-items: center;
`;

const button = css`
  background-color: #007bff;
  border-color: #007bff;
  border-radius: 0.25rem;
  color: #fff;
  cursor: pointer;
  font-size: 1rem;
  line-height: 1.5;
  padding: 0.375rem 0.75rem;
`;

const EventDetail = (props) => {
  const { id, startAt, available } = props;

  const [checked, setChecked] = useState(available);

  const onChecked = (checked) => {
    setChecked(checked);
  };

  const onClick = () => {
    console.log(id);
  };

  return (
    <div css={style}>
      <p>Día: {moment(startAt).locale("es").format("LLLL")}</p>
      <Switch
        onChange={onChecked}
        checked={checked}
        checkedIcon={false}
        uncheckedIcon={false}
      />
      <button css={button} onClick={onClick}>
        Bookear
      </button>
    </div>
  );
};

export default EventDetail;
