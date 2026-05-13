import React, { forwardRef } from 'react';
import { Button } from 'antd';

const mapButtonType = (color, basic) => {
  if (basic) return 'default';
  if (color === 'blue') return 'primary';
  return 'default';
};

const AppButton = forwardRef(function AppButton(
  {
    className = '',
    children,
    color,
    basic,
    danger,
    fluid = false,
    ...props
  },
  ref,
) {
  const nextClassName = ['router-ui-button', fluid ? 'fluid' : '', className]
    .filter(Boolean)
    .join(' ');
  const nextDanger = danger || color === 'red';
  return (
    <Button
      {...props}
      ref={ref}
      className={nextClassName}
      type={mapButtonType(color, basic)}
      danger={nextDanger}
    >
      {children}
    </Button>
  );
});

export default AppButton;
