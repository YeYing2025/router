import React from 'react';
import { Drawer } from 'antd';

function AppDrawer({ className = '', children, open, onClose, title, ...props }) {
  const nextClassName = ['router-ui-drawer', className].filter(Boolean).join(' ');
  return (
    <Drawer
      {...props}
      className={nextClassName}
      open={open}
      onClose={onClose}
      title={title}
    >
      {children}
    </Drawer>
  );
}

export default AppDrawer;
