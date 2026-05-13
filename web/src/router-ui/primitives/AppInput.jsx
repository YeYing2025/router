import React from 'react';
import { Input, Space } from 'antd';

function AppInput({
  className = '',
  value,
  onChange,
  name,
  disabled,
  readOnly,
  placeholder,
  type,
  maxLength,
  fluid = false,
  icon,
  iconPosition,
  action,
  loading,
  ...props
}) {
  const nextClassName = ['router-ui-input', fluid ? 'fluid' : '', className]
    .filter(Boolean)
    .join(' ');
  const prefix =
    icon && iconPosition !== 'right' ? <i className={`${icon} icon`} /> : undefined;
  const suffix =
    icon && iconPosition === 'right' ? <i className={`${icon} icon`} /> : undefined;

  const handleChange = (event) => {
    if (typeof onChange === 'function') {
      onChange(event, {
        name,
        value: event?.target?.value ?? '',
      });
    }
  };

  const inputNode = (
    <Input
      {...props}
      className={nextClassName}
      value={value}
      onChange={handleChange}
      disabled={disabled}
      readOnly={readOnly}
      placeholder={placeholder}
      type={type}
      maxLength={maxLength}
      prefix={prefix}
      suffix={suffix}
    />
  );

  if (!action) {
    return inputNode;
  }

  return (
    <Space.Compact className={fluid ? 'fluid' : undefined} block={fluid}>
      {inputNode}
      {action}
    </Space.Compact>
  );
}

export default AppInput;
