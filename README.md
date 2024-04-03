<div align=center>

# warp-plus

## Automated WARP+ referrals

</div>

## Usage

-   Get an ID by registering on the app and copying the ID field under Settings > Advanced > Diagnostics > Client Configuration
-   ```
    docker build -t warp-plus .
    ```
-   ```
    docker run -d -e ID="{ID}" -e ERROR_INTERVAL=10 -e SUCCESS_INTERVAL=60 --name warp-plus warp-plus
    ```

## Credits

[nxvvvv/warp-plus](https://github.com/nxvvvv/warp-plus)
