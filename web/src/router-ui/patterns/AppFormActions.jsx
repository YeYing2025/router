import React from 'react';

function AppFormActions({
  className = '',
  align = 'end',
  children,
}) {
  const nextClassName = [
    'router-ui-form-actions',
    align === 'start' ? 'align-start' : '',
    align === 'center' ? 'align-center' : '',
    className,
  ]
    .filter(Boolean)
    .join(' ');

  return <div className={nextClassName}>{children}</div>;
}

export default AppFormActions;
