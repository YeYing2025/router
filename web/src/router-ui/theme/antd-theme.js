import { routerTokens } from './tokens';

export const antdTheme = {
  token: {
    colorPrimary: routerTokens.colorPrimary,
    colorSuccess: routerTokens.colorSuccess,
    colorWarning: routerTokens.colorWarning,
    colorError: routerTokens.colorError,
    colorText: routerTokens.colorText,
    colorTextSecondary: routerTokens.colorTextSecondary,
    colorBorder: routerTokens.colorBorder,
    colorBgContainer: routerTokens.colorBgContainer,
    borderRadius: routerTokens.borderRadius,
    borderRadiusSM: routerTokens.borderRadiusSM,
    borderRadiusLG: routerTokens.borderRadiusLG,
    controlHeight: routerTokens.controlHeight,
    controlHeightSM: routerTokens.controlHeightSM,
    controlHeightLG: routerTokens.controlHeightLG,
    fontFamily: routerTokens.fontFamily,
    fontSize: routerTokens.fontSize,
  },
  components: {
    Button: {
      borderRadius: routerTokens.borderRadius,
      controlHeight: routerTokens.controlHeight,
      fontWeight: 500,
    },
    Descriptions: {
      labelBg: '#fafafa',
      titleMarginBottom: routerTokens.spaceSM,
    },
    Input: {
      activeBorderColor: routerTokens.colorPrimary,
      hoverBorderColor: routerTokens.colorPrimary,
    },
    Select: {
      optionSelectedBg: '#e6f4ff',
      optionActiveBg: '#f5f6f7',
    },
    Segmented: {
      trackBg: '#f5f6f7',
      itemColor: '#667085',
      itemHoverColor: '#1f2329',
      itemHoverBg: '#ffffff',
      itemSelectedBg: '#ffffff',
      itemSelectedColor: routerTokens.colorPrimary,
      itemActiveBg: '#ffffff',
      trackPadding: 4,
      borderRadius: 10,
    },
    Switch: {
      colorPrimary: routerTokens.colorPrimary,
      colorPrimaryHover: '#4096ff',
      handleBg: '#ffffff',
      trackHeight: 22,
      trackMinWidth: 42,
      trackPadding: 2,
    },
    Modal: {
      borderRadiusLG: routerTokens.borderRadiusLG,
    },
    Layout: {
      siderBg: '#f8fafc',
      triggerBg: '#f8fafc',
      triggerColor: routerTokens.colorText,
    },
    Menu: {
      itemBg: 'transparent',
      subMenuItemBg: 'transparent',
      itemColor: '#475569',
      itemHoverColor: '#1677ff',
      itemSelectedColor: '#1677ff',
      itemSelectedBg: '#eaf3ff',
      itemHoverBg: '#f1f5f9',
      activeBarWidth: 0,
      collapsedIconSize: 16,
      groupTitleColor: '#94a3b8',
    },
    Table: {
      headerBg: '#fafafa',
      rowHoverBg: '#f8fbff',
    },
    Tabs: {
      itemSelectedColor: routerTokens.colorPrimary,
      itemHoverColor: routerTokens.colorPrimary,
      inkBarColor: routerTokens.colorPrimary,
    },
  },
};

export default antdTheme;
