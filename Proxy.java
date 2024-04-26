// Definir la interface Artwork
interface Artwork{
    display();
}

// ConcreteArtwork representa artworks individualmente
class ConcreteArtwork implements Artwork{
    name;
    image; 

    constructor(name, image){
        this.name = name;
        this.image = image;
    }
    display(){
        print("Displaying " + this.name);
    }
}

// Definir la clase proxy para Artwork
class ArtworkProxy implements Artwork{
    artwork;
    highResolutionImage; 

    constructor(name){
        this.artwork = new ConcreteArtwork(name, null);
    }
    display(){
        if (!this.highResolutionImage){
            this.highResolutionImage = loadHighResolutionImage(this.artwork.name);
        }
        this.artwork.display();
    }
}

// ArtGallery maneja los artworks
class ArtGallery{
    artworks;
    constructor(){
        this.artworks = [];
    }
    addArtwork(artwork){
        this.artworks.push(artwork);
    }
    displayAllArtworks(){
        for each (artwork in this.artworks){
            artwork.display();
        }
    }
}

// Ejemplo de uso ---------
gallery = new ArtGallery();
// Crear y agregar Artwork a gallery usando proxies.
artwork1 = new ArtworkProxy("Artwork1");
artwork2 = new ArtworkProxy("Artwork2");
gallery.addArtwork(artwork1);
gallery.addArtwork(artwork2);
// Mostrar Artwork de gallery
gallery.displayAllArtworks();
