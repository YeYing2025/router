import React from 'react';
import { Segmented } from 'antd';

const normalizeOptions = (options) =>
  (Array.isArray(options) ? options : []).map((option) => ({
    label: option?.label ?? option?.text ?? option?.name ?? option?.value,
    value: option?.value,
    disabled: option?.disabled === true,
  }));

function AppSegmented({
  className = '',
  options = [],
  value,
  name,
  block = false,
  disabled = false,
  onChange,
  ...props
}) {
  const nextClassName = ['router-ui-segmented', className]
    .filter(Boolean)
    .join(' ');

  const handleChange = (nextValue) => {
    if (typeof onChange === 'function') {
      onChange(null, {
        name,
        value: nextValue,
      });
    }
  };

  return (
    <Segmented
      {...props}
      className={nextClassName}
      options={normalizeOptions(options)}
      value={value}
      block={block}
      disabled={disabled}
      onChange={handleChange}
    />
  );
}

export default AppSegmented;
