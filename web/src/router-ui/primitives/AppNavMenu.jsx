import React from 'react';
import { Menu } from 'antd';

function AppNavMenu({ className = '', items = [], ...props }) {
  const nextClassName = ['router-ui-nav-menu', className].filter(Boolean).join(' ');
  return <Menu {...props} className={nextClassName} items={items} />;
}

export default AppNavMenu;
