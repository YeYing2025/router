import React from 'react';
import { Input } from 'antd';

function AppTextarea({
  className = '',
  value,
  onChange,
  name,
  disabled,
  readOnly,
  placeholder,
  rows,
  maxLength,
  autoSize,
  ...props
}) {
  const nextClassName = ['router-ui-textarea', className]
    .filter(Boolean)
    .join(' ');

  const handleChange = (event) => {
    if (typeof onChange === 'function') {
      onChange(event, {
        name,
        value: event?.target?.value ?? '',
      });
    }
  };

  return (
    <Input.TextArea
      {...props}
      className={nextClassName}
      value={value}
      onChange={handleChange}
      disabled={disabled}
      readOnly={readOnly}
      placeholder={placeholder}
      rows={rows}
      maxLength={maxLength}
      autoSize={autoSize}
    />
  );
}

export default AppTextarea;
