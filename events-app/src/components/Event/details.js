/** @jsx jsx */
import { css, jsx } from "@emotion/core";
import Switch from "react-switch";
import { useState } from "react";
import moment from "moment";

const style = css`
  display: flex;
  width: 100%;
  justify-content: space-around;
`;

const EventDetail = (props) => {
  const { startAt, available } = props;

  const [checked, setChecked] = useState(available);

  const handleChecked = (checked) => {
    setChecked(checked);
  };

  return (
    <div css={style}>
      <p>Día: {moment(startAt).locale("es").format("LLLL")}</p>
      <Switch
        onChange={handleChecked}
        checked={checked}
        checkedIcon={false}
        uncheckedIcon={false}
      />
    </div>
  );
};

export default EventDetail;
