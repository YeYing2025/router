import React from 'react';
import AppButton from './AppButton';
import AppIcon from './AppIcon';
import AppTooltip from './AppTooltip';

function AppTableActionButton({
  title,
  icon,
  className = '',
  color,
  disabled = false,
  ...props
}) {
  const button = (
    <AppButton
      {...props}
      type='button'
      color={color}
      disabled={disabled}
      aria-label={title}
      className={['router-inline-button', 'router-table-action-button', className]
        .filter(Boolean)
        .join(' ')}
    >
      <AppIcon name={icon} />
    </AppButton>
  );

  return (
    <AppTooltip title={title}>
      <span className='router-table-action-button-wrap'>{button}</span>
    </AppTooltip>
  );
}

export default AppTableActionButton;
