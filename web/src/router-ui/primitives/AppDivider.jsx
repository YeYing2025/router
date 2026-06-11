import React from 'react';
import { Divider } from 'antd';

function AppDivider({
  className = '',
  children,
  horizontal = false,
  orientation,
  titlePlacement,
  ...props
}) {
  const resolvedTitlePlacement = titlePlacement || (horizontal ? 'center' : orientation);
  return (
    <Divider
      {...props}
      className={className}
      titlePlacement={resolvedTitlePlacement}
    >
      {children}
    </Divider>
  );
}

export default AppDivider;
