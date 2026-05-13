import React from 'react';
import { Tooltip } from 'antd';

function AppTooltip({ children, ...props }) {
  return <Tooltip {...props}>{children}</Tooltip>;
}

export default AppTooltip;
