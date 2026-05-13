import React from 'react';

function AppFormRow({ className = '', children }) {
  const nextClassName = ['router-ui-form-row', className]
    .filter(Boolean)
    .join(' ');
  return <div className={nextClassName}>{children}</div>;
}

export default AppFormRow;
