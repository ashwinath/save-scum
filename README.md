# Save Scum

Backup your data like a save scum would.

## Sample yaml

```yaml
Files:
- From: /home/ashwin/Downloads
  To: /home/ashwin/tmp-rsync
  Flags:
  - '--recursive'
  - '--progress'
  - '--checksum'
  - '--owner'
  - '--group'
  Chown:
    Enabled: True
    User: ashwin
    Group: ashwin
```
