import React, { forwardRef } from 'react';
import { InputNumber } from 'antd';

const AppInputNumber = forwardRef(function AppInputNumber(
  {
    className = '',
    fluid = false,
    name,
    onChange,
    ...props
  },
  ref,
) {
  const nextClassName = ['router-ui-input-number', fluid ? 'fluid' : '', className]
    .filter(Boolean)
    .join(' ');

  const handleChange = (value) => {
    if (typeof onChange === 'function') {
      onChange(null, {
        name,
        value,
      });
    }
  };

  return (
    <InputNumber
      {...props}
      ref={ref}
      className={nextClassName}
      onChange={handleChange}
    />
  );
});

export default AppInputNumber;
