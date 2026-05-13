import React from 'react';
import { Dropdown } from 'antd';

function AppMenuDropdown({
  className = '',
  menuClassName = '',
  items = [],
  disabled = false,
  children,
  trigger = ['click'],
  placement = 'bottomLeft',
}) {
  const nextClassName = ['router-ui-menu-dropdown', className]
    .filter(Boolean)
    .join(' ');
  const nextMenuClassName = ['router-ui-menu-overlay', menuClassName]
    .filter(Boolean)
    .join(' ');

  const menuItems = items.map((item) => ({
    key: item.key,
    disabled: item.disabled === true,
    danger: item.danger === true,
    label: (
      <span
        className={[
          'router-ui-menu-item-label',
          item.active ? 'active' : '',
          item.className || '',
        ]
          .filter(Boolean)
          .join(' ')}
      >
        {item.icon ? <span className='router-ui-menu-item-icon'>{item.icon}</span> : null}
        <span>{item.label}</span>
      </span>
    ),
    onClick: item.onClick,
  }));

  return (
    <Dropdown
      disabled={disabled}
      trigger={trigger}
      placement={placement}
      menu={{ items: menuItems }}
      classNames={{ root: nextMenuClassName }}
    >
      <span className={nextClassName}>{children}</span>
    </Dropdown>
  );
}

export default AppMenuDropdown;
