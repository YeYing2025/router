import React from 'react';
import { Divider } from 'antd';

function AppDivider({
  className = '',
  children,
  horizontal = false,
  ...props
}) {
  return (
    <Divider
      {...props}
      className={className}
      orientation={horizontal ? 'center' : props.orientation}
    >
      {children}
    </Divider>
  );
}

export default AppDivider;
