import React from 'react';
import { Popover } from 'antd';

function AppPopover({ children, content, trigger, ...props }) {
  const triggerNode = children ?? trigger;
  return (
    <Popover {...props} content={content}>
      {triggerNode}
    </Popover>
  );
}

export default AppPopover;
