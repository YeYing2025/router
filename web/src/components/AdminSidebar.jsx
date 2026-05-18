import React, { useEffect, useMemo, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import {
  ADMIN_MENU_GROUPS,
  isAdminRouteActive,
} from '../constants/adminMenu';
import { AppIcon, AppNavMenu } from '../router-ui';

const SIDEBAR_GROUP_OPEN_STORAGE_KEY = 'router_admin_sidebar_group_open_v2';

const buildDefaultOpenKeys = () => ADMIN_MENU_GROUPS.map((group) => group.key);

const buildInitialOpenKeys = () => {
  const defaults = buildDefaultOpenKeys();
  if (typeof window === 'undefined') {
    return defaults;
  }
  const raw = (localStorage.getItem(SIDEBAR_GROUP_OPEN_STORAGE_KEY) || '').trim();
  if (raw === '') {
    return defaults;
  }
  try {
    const parsed = JSON.parse(raw);
    if (!Array.isArray(parsed)) {
      return defaults;
    }
    const allowed = new Set(defaults);
    return parsed.filter((key) => allowed.has(key));
  } catch {
    return defaults;
  }
};

const AdminSidebar = ({ compact = false }) => {
  const { t } = useTranslation();
  const location = useLocation();
  const navigate = useNavigate();
  const [openKeys, setOpenKeys] = useState(buildInitialOpenKeys);

  const selectedKeys = useMemo(() => {
    const active = [];
    ADMIN_MENU_GROUPS.forEach((group) => {
      group.items.forEach((item) => {
        if (isAdminRouteActive(location, item.to)) {
          active.push(item.to);
        }
      });
    });
    return active;
  }, [location]);

  useEffect(() => {
    if (typeof window === 'undefined') {
      return;
    }
    localStorage.setItem(SIDEBAR_GROUP_OPEN_STORAGE_KEY, JSON.stringify(openKeys));
  }, [openKeys]);

  useEffect(() => {
    if (compact || selectedKeys.length === 0) {
      return;
    }
    const activeGroupKeys = ADMIN_MENU_GROUPS.filter((group) =>
      group.items.some((item) => selectedKeys.includes(item.to)),
    ).map((group) => group.key);
    if (activeGroupKeys.length === 0) {
      return;
    }
    setOpenKeys((previous) => {
      const next = Array.from(new Set([...previous, ...activeGroupKeys]));
      return next.length === previous.length &&
        next.every((item, index) => item === previous[index])
        ? previous
        : next;
    });
  }, [compact, selectedKeys]);

  const items = useMemo(
    () =>
      ADMIN_MENU_GROUPS.map((group) => ({
        key: group.key,
        icon: <AppIcon name={group.icon} />,
        label: t(group.name),
        children: group.items.map((item) => ({
          key: item.to,
          icon: <AppIcon name={item.icon} />,
          label: t(item.name),
        })),
      })),
    [t],
  );

  return (
    <AppNavMenu
      className='router-admin-nav-menu'
      mode='inline'
      inlineCollapsed={compact}
      triggerSubMenuAction={compact ? 'click' : 'hover'}
      items={items}
      selectedKeys={selectedKeys}
      {...(!compact
        ? {
            openKeys,
            onOpenChange: (nextKeys) => setOpenKeys(nextKeys),
          }
        : {})}
      onClick={({ key }) => {
        if (typeof key === 'string' && key.startsWith('/')) {
          navigate(key);
        }
      }}
    />
  );
};

export default AdminSidebar;
