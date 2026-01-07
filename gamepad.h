struct libevdev_uinput *create_controller();
void destroy_device(struct libevdev_uinput *uidev);
void btn_down(struct libevdev_uinput *uidev, unsigned int code);
void syn(struct libevdev_uinput *uidev);
void btn_up(struct libevdev_uinput *uidev, unsigned int code);