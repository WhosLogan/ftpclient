export namespace ftpclient {
	
	export class FtpEntry {
	    Perms: string;
	    Links: number;
	    UID: number;
	    GID: number;
	    Size: number;
	    Month: string;
	    Day: number;
	    Time: string;
	    Name: string;
	    IsDir: boolean;
	
	    static createFrom(source: any = {}) {
	        return new FtpEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Perms = source["Perms"];
	        this.Links = source["Links"];
	        this.UID = source["UID"];
	        this.GID = source["GID"];
	        this.Size = source["Size"];
	        this.Month = source["Month"];
	        this.Day = source["Day"];
	        this.Time = source["Time"];
	        this.Name = source["Name"];
	        this.IsDir = source["IsDir"];
	    }
	}

}

