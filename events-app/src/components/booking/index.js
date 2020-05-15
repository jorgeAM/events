/** @jsx jsx */
import { css, jsx } from "@emotion/core";
import { useState } from "react";

const container = css`
  padding: 0 10%;

  > h1 {
    margin: 0;
    margin-bottom: 10px;
    text-align: center;
  }
`;

const innerContainer = css`
  display: flex;
  align-items: center;
  margin-bottom: 16px;

  > label {
    margin-right: 24px;
  }
`;

const input = css`
  padding: 8px;
  border-radius: 5px;
  width: 100%;
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

const Booking = () => {
  const [seats, setSeats] = useState(0);

  const onChange = (ev) => {
    setSeats(ev.target.value);
  };

  const onClick = () => {
    console.log(seats);
  };

  return (
    <div css={container}>
      <h1>BOOKING</h1>
      <div css={innerContainer}>
        <label>Asientos: </label>
        <input
          css={input}
          type="number"
          min={0}
          value={seats}
          onChange={onChange}
        />
      </div>
      <button css={button} onClick={onClick}>
        Guardar
      </button>
    </div>
  );
};

export default Booking;
